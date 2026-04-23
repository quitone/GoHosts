import { conf } from '@wailsjs/go/models'
import { GetSystemHosts, GetSchemes } from '@wailsjs/go/main/App'

export const useHostsStore = defineStore('hosts', () => {
  const hostsData = ref<conf.Scheme>({
    id: '',
    name: '',
    type: '',
    enabled: false,
    content: '',
    createdAt: '',
    updatedAt: '',
    refreshTime: 0,
  })
  const schemes = ref<conf.Scheme[]>([])

  const curHostsContent = ref<string>('')
  let systemHostsContent = ''
  const isHostsReadonly = ref<boolean>(false)
  const hostsId = ref<string>('')

  GetSchemes().then(data => {
    schemes.value = data
    console.log(data, 'data')
  })
  const openHostsConfig = (id: string) => {}
  const loadSystemHosts = () => {
    GetSystemHosts().then(data => {
      systemHostsContent = data
      isHostsReadonly.value = true
    })
  }
  const showSystemHosts = () => {
    curHostsContent.value = systemHostsContent
    isHostsReadonly.value = true
  }

  return {
    hostsData,
    curHostsContent,
    isHostsReadonly,
    hostsId,

    loadSystemHosts,
    showSystemHosts,
  }
})
