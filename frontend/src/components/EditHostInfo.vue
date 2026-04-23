<template>
  <n-drawer :show="appStore.editSchemaVisible" mask-closable close-on-esc @hide="appStore.editSchemaVisible = false" width="600px">
    <n-drawer-content>
      <template #header>
        <div class="text-left"><i class="i-mdi-library-edit-outline text-2xl mr-2"></i> 编辑Hosts</div>
      </template>

      <n-form class="text-left">
        <n-form-item label="Hosts 标题">
          <n-input v-model="formData.name" placeholder="" />
        </n-form-item>

        <n-form-item label="Hosts 类型">
          <n-radio-group :value="formData.type" @update:value="formData.type = $event">
            <n-radio value="0" class="mr-4">本地</n-radio>
            <n-radio value="1">远程</n-radio>
          </n-radio-group>
        </n-form-item>

        <template v-if="formData.type === '1'">
          <n-form-item label="Hosts URL">
            <n-input v-model="formData.url" placeholder="http:// 或 https:// " />
          </n-form-item>

          <n-form-item label="Hosts 刷新时间">
            <n-select v-if="!isCustomTime" v-model:value="formData.refreshTime" :options="refreshTimeOptions" />
            <div v-else class="grid max-w-xl grid-cols-4 gap-2">
              <n-input-number v-model:value="customTimeArr[0]" size="small" :min="0" :max="99">
                <template #suffix>天</template>
              </n-input-number>
              <n-input-number v-model:value="customTimeArr[1]" size="small" :min="0" :max="23">
                <template #suffix>时</template>
              </n-input-number>
              <n-input-number v-model:value="customTimeArr[2]" size="small" :min="0" :max="59">
                <template #suffix>分</template>
              </n-input-number>
              <n-input-number v-model:value="customTimeArr[3]" size="small" :min="0" :max="59">
                <template #suffix>秒</template>
              </n-input-number>
            </div>

            <!-- <n-switch v-model:value="isCustomTime" class="ml-4" size="small" /> -->
            <n-checkbox v-model:checked="isCustomTime" label="自定义时间" class="w-138px ml-4" />
          </n-form-item>
        </template>
      </n-form>

      <template #footer>
        <n-button class="mr-4" type="primary" @click="handleSave">保存</n-button>
        <n-button @click="handleCancel" type="default">取消</n-button>
      </template>
    </n-drawer-content>
  </n-drawer>
</template>

<script setup lang="ts">
import { useAppStore } from '@/stores/app.store'
import type { conf } from '@wailsjs/go/models'

const appStore = useAppStore()
const formData = reactive<Pick<conf.Scheme, 'name' | 'type' | 'url' | 'refreshTime' | 'id'>>({
  id: '',
  name: '',
  type: '0',
  refreshTime: 0,
  url: '',
})

const isCustomTime = ref(true)
const customTimeArr = ref<number[]>([0, 0, 0, 0])
const refreshTimeOptions = [
  { value: 0, label: '从不' },
  { value: 60, label: '1 分钟' },
  { value: 60 * 5, label: '5 分钟' },
  { value: 60 * 15, label: '15 分钟' },
  { value: 60 * 60, label: '1 小时' },
  { value: 60 * 60 * 12, label: '12 小时' },
  { value: 60 * 60 * 24, label: '1 天' },
  { value: 60 * 60 * 24 * 7, label: '7 天' },
]

function handleSave() {
  appStore.editSchemaVisible = false
}
const handleCancel = () => {
  appStore.editSchemaVisible = false
}
</script>

<style>
.n-form-item-label__text {
  font-size: 16px;
  font-weight: 600;
  line-height: 24rpx;
}
</style>
