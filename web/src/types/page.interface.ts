export interface Page {
  id: string
  title: string
  content: string
  file_list: string
  author: string
  authorIP: string
  updated_ts: number
  created_ts: number
}

export interface Post {
  msg: string
}
