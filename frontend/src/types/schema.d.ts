export interface Scheme {
  id: string
  name: string
  type: string
  content?: string
  url?: string
  enabled: boolean
  createdAt: string // ISO 8601 format string from Go time.Time
  updatedAt: string // ISO 8601 format string from Go time.Time
}

export interface Config {
  schemes: Scheme[]
  activeScheme: string
}

declare global {
  interface Window {
    go: {
      app: {
        App: any
      }
    }
  }
}