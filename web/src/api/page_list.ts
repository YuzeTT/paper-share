import axios from 'axios'
import type { Page } from '~/types/page.interface'

const getPages = async (): Promise<Page[] | null> => {
  const res: Page[] | null = await axios
    .get('http://localhost:3000/api/pages')
    .then((res) => {
      return res.data
    })
    .catch((err) => {
      console.error(err)
      return null
    })

  return res
}

export default getPages
