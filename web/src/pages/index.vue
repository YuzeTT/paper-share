<script setup lang="ts" generic="T extends any, O extends any">
import getPages from '~/api/page_list'
import type { Page } from '~/types/page.interface'

defineOptions({
  name: 'IndexPage',
})

let pages = $ref<Page[] | null>()

// axios.get('http://localhost:3000/api/pages').then(data => {
//   pages = data
// })
onMounted(async () => {
  pages = await getPages()
})
</script>

<template>
  <div>
    <div v-if="!pages">
      loading
    </div>
    <div v-else>
      <div v-for="(item, i) in pages" :key="i">
        {{ item }}
      </div>
    </div>
  </div>
</template>
