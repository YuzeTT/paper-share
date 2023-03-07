import axios from 'axios'
import type { PagesResponse } from '~/interface/page.interface'

const getPages = async (): Promise<PagesResponse> => {
  const res: PagesResponse = await axios
    .get('http://localhost:3000/api/pages')
    .then((res) => {
      return res.data as PagesResponse
    })
    .catch((err) => {
      console.error(err)
      return err.response
        ? err.response.data
        : { message: 'server error' }
    })

  return res
}

export default getPages
