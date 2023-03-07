<script setup lang="ts" generic="T extends any, O extends any">
import getPages from '~/api/page'
import type { PagesResponse } from '~/interface/page.interface'

defineOptions({
  name: 'IndexPage',
})

let pages = $ref<PagesResponse>()

// axios.get('http://localhost:3000/api/pages').then(data => {
//   pages = data
// })
onMounted(async () => {
  const res: PagesResponse = await getPages()
  if (res.message !== 'server error')
    pages = res
})
</script>

<template>
  <div>
    <div v-if="!pages">
      loading
    </div>
    <div v-else-if="pages.message === 'server error'">
      boom!
    </div>
    <div v-else>
      <div v-for="(item, i) in pages" :key="i">
        {{ item }}
      </div>
    </div>
  </div>
</template>
