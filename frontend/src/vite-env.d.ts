/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

declare module '@wailsjs/go/main/App' {
  import { runtime } from '@wailsjs/go/models'
  export function LoadConfig(): Promise<any>
  export function SaveConfig(config: any): Promise<void>
  export function GetSchemes(): Promise<any[]>
  export function SaveScheme(scheme: any): Promise<void>
  export function DeleteScheme(id: string): Promise<void>
  export function GetSystemHosts(): Promise<string>
  export function GetEnv(): Promise<runtime.EnvironmentInfo>
  export function Quit(): Promise<void>
  export function SendNotification(title: string, body: string): Promise<void>
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
      refreshTime: number
    }
  }
}
