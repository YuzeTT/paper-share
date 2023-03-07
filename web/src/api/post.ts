import axios from 'axios'
import type { Page } from '~/types/page.interface'

const postPage = async (title: string, author: string, content: string): Promise<Page> => {
  const res: Page = await axios
    .post('http://localhost:3000/api/post', { title, author, content })
    .then((res) => {
      return res.data
    })
    .catch((err) => {
      console.error(err)
      return {
        msg: '请求失败',
      }
    })

  return res
}

export default postPage
