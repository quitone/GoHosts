import { GetEnv } from '@wailsjs/go/main/App'

export async function getPlatform(): Promise<{
  platform: 'windows' | 'macos' | 'linux'
  arch: 'amd64' | 'arm64'
  buildType: 'production' | 'debug' | 'dev'
}> {
  const envInfo = await GetEnv()
  // const envInfo = {
  //   platform: 'windows',
  //   arch: 'amd64',
  //   buildType: 'production',
  // } as const
  console.log(`当前平台: ${envInfo.platform}`)
  return envInfo
}

export function debounce(fn: Function, delay = 300, immediate = true) {
  let timeoutId: any
  if (immediate) {
    return function (...args: any[]) {
      if (timeoutId) return

      fn.apply(null, args)
      timeoutId = setTimeout(() => {
        timeoutId = null
      }, delay)
    }
  } else {
    return function (...args: any[]) {
      clearTimeout(timeoutId)
      timeoutId = setTimeout(() => fn.apply(null, args), delay)
    }
  }
}
