/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

declare module '@wailsjs/go/app/App' {
  export function LoadConfig(): Promise<any>
  export function SaveConfig(config: any): Promise<void>
  export function GetSchemes(): Promise<any[]>
  export function SaveScheme(scheme: any): Promise<void>
  export function DeleteScheme(id: string): Promise<void>
  export function GetSystemHosts(): Promise<string>
}

declare module '@wailsjs/go/models' {
  export namespace conf {
    export class Config {
      schemes: any[]
      activeScheme: string
    }
    export class Scheme {
      id: string
      name: string
      type: string
      content?: string
      url?: string
      enabled: boolean
      createdAt: any
      updatedAt: any
    }
  }
}
