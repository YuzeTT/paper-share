import axios from 'axios'
import type { Page } from '~/types/page.interface'

const getPage = async (id_: string): Promise<Page> => {
  const res: Page = await axios
    .get('http://localhost:3000/api/page', { params: { id: id_ } })
    .then((res) => {
      return res.data
    })
    .catch((err) => {
      console.error(err)
      return {
        id: '404',
        title: '404 - 未找到页面',
        content: '',
        file_list: '{}',
        author: '系统',
        author_ip: '0',
        created_ts: 0,
        updated_ts: 0,
      }
    })

  return res
}

export default getPage
