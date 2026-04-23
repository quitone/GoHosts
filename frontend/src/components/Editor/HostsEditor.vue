<template>
  <div class="relative container">
    <!-- <div ref="editorRef" class="editor-ref"></div> -->
    <div ref="editor" tabindex="0" class="editor"></div>
  </div>
</template>
<!-- v-model="hostsStore.curHostsContent" -->
<!-- <n-tabs v-model:value="activeTab" type="line" class="flex-1 flex flex-col">
      <n-tab-pane name="editor" tab="方案编辑">
        <div v-if="editingScheme" class="flex flex-col h-full px-16px pb-16px">
          <div class="flex-between gap-12px mb-12px flex-shrink-0">
            <n-input v-model:value="editingScheme.name" placeholder="方案名称" class="max-w-300px" />
            <div class="toolbar-actions">
              <n-button :loading="store.saving" type="primary" @click="handleSaveScheme"> 保存 </n-button>
            </div>
          </div>
          <div class="flex-1 overflow-hidden min-h-0">
            <textarea
              v-model="editingScheme.content"
              class="w-full h-full bg-neutral-800 text-neutral-100 b b-neutral-700 rounded-6px p-12px font-mono text-14px leading-1.6 resize-none outline-none box-border focus:b-primary"
              placeholder="# 在此输入 hosts 内容&#10;127.0.0.1 localhost"
              spellcheck="false" />
          </div>
          <div class="flex-between items-center mt-12px text-13px flex-shrink-0">
            <span class="text-neutral-400">{{ store.saving ? '保存中...' : '就绪' }}</span>
            <span v-if="store.error" class="text-error">{{ store.error }}</span>
          </div>
        </div>
        <div v-else class="flex-center h-full text-neutral-400 text-16px">
          <p>请从左侧选择一个方案进行编辑</p>
        </div>
      </n-tab-pane>

      <n-tab-pane name="system" tab="系统 Hosts">
        <HostsView />
      </n-tab-pane>
    </n-tabs> -->

<script setup lang="tsx">
import { useHostsStore } from '@/stores/hosts.store'
import type { Scheme } from '@/types/schema'
import { withLineNumbers } from 'codejar-linenumbers'
import { CodeJar, type Position } from 'codejar'
import { highlightHosts, toggleCommentByLine, toggleCommentBySelection } from './hosts-highlight'
import { debounce } from '@/utils/utils'
import { normalizeLineEndings } from '@/constant/newlines'

// CodeJar(editor, withLineNumbers(highlight), {tab: '\t'})

const hostsStore = useHostsStore()
const { isHostsReadonly, hostsId, curHostsContent } = storeToRefs(hostsStore)
const editor = ref<HTMLDivElement>()
// const editorRef = ref<HTMLDivElement>()

let jar: ReturnType<typeof CodeJar>
const onEditorUpdate = (code: string) => {
  hostsStore.curHostsContent = code
}
onMounted(() => {
  jar = CodeJar(
    editor.value!,
    withLineNumbers(highlightHosts, {
      class: 'absolute lh-6',
      width: '32px',
      backgroundColor: 'var(--editor-gutter-bg)',
      color: 'var(--editor-line-number-color)',
    }),
  )
  jar.updateCode(unref(curHostsContent), false)

  setContentEditable(unref(isHostsReadonly))
  const onMountClick = (event: MouseEvent) => {
    const target = event.target as HTMLElement | null
    const gutter = target?.closest('.codejar-linenumbers')
    if (!gutter) return

    const lineHeight = parseFloat(window.getComputedStyle(unref(editor)!).lineHeight) || 24
    const scrollContainer = gutter.closest('.codejar-wrap') ?? unref(editor)
    const relativeY = event.clientY - gutter.getBoundingClientRect().top + scrollContainer!.scrollTop
    const lineCount = Math.max(1, jar.toString().split('\n').length)
    const lineIndex = Math.max(0, Math.min(lineCount - 1, Math.floor(relativeY / lineHeight)))

    event.preventDefault()
    onGutterClick(lineIndex)
  }

  jar.onUpdate(onEditorUpdate)
  jar.updateCode('', false)
  editor.value!.addEventListener('click', onMountClick)
  // loadContent(hosts_id).catch((e) => console.error(e))
})
const setContentEditable = (isReadonly: boolean) => {
  const editorInst = unref(editor)
  if (!editorInst) return
  editorInst.setAttribute('contenteditable', !isReadonly ? 'false' : 'plaintext-only')
  editorInst.setAttribute('aria-readonly', !isReadonly ? 'true' : 'false')
}
watch(curHostsContent, code => {
  jar && jar.updateCode(code, false)
})
watch(isHostsReadonly, setContentEditable)
const editingScheme = ref<Scheme | null>(null)

// useEffect(() => {
//     ref_hosts_id.current = hosts_id
//   }, [hosts_id])

//   useEffect(() => {
//     ref_is_read_only.current = is_read_only
//   }, [is_read_only])

let pendingFindTimer: number | null = null
const clearPendingFind = () => {
  if (pendingFindTimer) {
    window.clearTimeout(pendingFindTimer)
    pendingFindTimer = null
  }
  pendingFindTimer = null
}

//   useEffect(() => clearPendingFind, [])

const handleSave = debounce((id: string, nextContent: string) => {
  console.log(id, nextContent)
  // actions
  //   .setHostsContent(id, nextContent)
  //   .then(() => agent.broadcast(events.hosts_content_changed, id))
  //   .catch((e) => console.error(e))
}, 1000)

/** Scroll the current selection/cursor into view after programmatic focus changes. */
const scrollSelectionIntoView = () => {
  if (!unref(editor)) return

  const selection = window.getSelection()
  if (!selection || selection.rangeCount === 0) return

  const range = selection.getRangeAt(0)
  const startNode = range.startContainer
  const target = startNode.nodeType === Node.TEXT_NODE ? startNode.parentElement : (startNode as Element | null)

  ;(target ?? unref(editor)!).scrollIntoView({
    block: 'nearest',
    inline: 'nearest',
  })
}

/** Restore a character-offset selection in the editor (used by find/show-source). */
const setSelection = (params: any) => {
  if (!jar || !unref(editor)) return
  const editorContent = jar.toString()

  const start = Math.max(0, Math.min(params.start, editorContent.length))
  const end = Math.max(0, Math.min(params.end, editorContent.length))
  jar.restore({ start, end, dir: '->' })
  editor.value!.focus()
  window.requestAnimationFrame(scrollSelectionIntoView)
}

/** Fetch and display the hosts content. Applies any pending find selection after loading. */
const loadContent = async (targetHostsId: string) => {
  if (!jar) return

  // const nextContent = normalizeLineEndings(
  //   targetHostsId === '0'
  //     ? await actions.getSystemHosts()
  //     : await actions.getHostsContent(targetHostsId),
  // )

  // if (ref_hosts_id.current !== targetHostsId) return

  // setContent(nextContent)
  // jar.updateCode(nextContent, false)

  // const pendingFind = ref_pending_find.current
  // if (pendingFind && pendingFind.item_id === targetHostsId) {
  //   setSelection(pendingFind)
  //   clearPendingFind()
  // }
}

// Position
const getCurrentSelection = (): Position => {
  // const jar = ref_jar.current
  // const editor = ref_editor.current
  const fallbackOffset = jar?.toString().length ?? 0
  if (!jar || !unref(editor)) {
    return { start: fallbackOffset, end: fallbackOffset, dir: '->' }
  }

  try {
    return jar.save()
  } catch {
    return { start: fallbackOffset, end: fallbackOffset, dir: '->' }
  }
}

const onChange = (nextContent: string) => {
  const normalizedContent = normalizeLineEndings(nextContent)
  curHostsContent.value = normalizedContent
  handleSave(hostsId, normalizedContent)
}

/** Push a programmatic edit into CodeJar: update content, restore selection, and record undo history. */
const applyEditorChange = (nextContent: string, nextSelection: Position) => {
  const editorInstance = unref(editor)
  if (!jar || !editorInstance) return

  editorInstance.focus()
  jar.recordHistory()
  jar.updateCode(nextContent, false)
  jar.restore(nextSelection)
  editorInstance.focus()
  jar.recordHistory()
  onChange(nextContent)
}

const toggleComment = () => {
  if (unref(isHostsReadonly)) return

  if (!jar) return

  const selection = getCurrentSelection()
  const next = toggleCommentBySelection(jar.toString(), selection.start, selection.end, true)
  if (!next.changed) return

  applyEditorChange(next.content, {
    start: next.selectionStart,
    end: next.selectionEnd,
    dir: '->',
  })
}

/** Handle a click on the line-number gutter to toggle comment on that line. */
const onGutterClick = (lineIndex: number) => {
  if (unref(isHostsReadonly)) return
  if (!jar) return

  const selection = getCurrentSelection()
  const next = toggleCommentByLine(jar.toString(), lineIndex, selection.start, selection.end)
  if (!next.changed) return

  applyEditorChange(next.content, {
    start: next.selectionStart,
    end: next.selectionEnd,
    dir: '->',
  })
}

//   useEffect(() => {
//     setEditorReadOnly(is_read_only)
//   }, [is_read_only])

//   useOnBroadcast(events.toggle_comment, toggleComment, [hosts_id])

//   useOnBroadcast(
//     events.show_source,
//     (params: IFindShowSourceParam) => {
//       if (params.item_id !== hosts_id || !ref_jar.current) {
//         clearPendingFind()
//         ref_pending_find.current = params
//         ref_pending_find_timer.current = window.setTimeout(clearPendingFind, 3000)
//         return
//       }

//       clearPendingFind()
//       setSelection(params)
//     },
//     [hosts_id],
//   )
</script>

<style lang="scss">
.container {
  height: 100%;
}

.editor {
  width: 100%;
  height: 100%;
  font-family: 'Fira Code', 'Monaco', monospace;
  font-size: 14px;
  padding-top: 16px;
  line-height: 24rpx;
  // background-color: #1e1e1e;

  color: #d4d4d4;
  overflow: hidden auto;
  outline: none;
  white-space: pre-wrap;
  tab-size: 2;
}

.codejar-wrap {
  .absolute {
    min-height: calc(100vh - var(--nav-bar-height));
  }
  height: calc(100vh - var(--nav-bar-height));
  overflow: hidden auto;
  display: flex;
  align-items: flex-start;
  .codejar-linenumbers-inner-wrap {
    max-width: 30px;
  }
}
</style>
