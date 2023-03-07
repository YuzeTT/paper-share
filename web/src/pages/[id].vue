<script setup lang="ts" generic="T extends any, O extends any">
import { marked } from 'marked'
import getPage from '~/api/page'
import type { Page } from '~/types/page.interface'
const props = defineProps<{ id: string }>()
defineOptions({
  name: 'Page',
})

let page = $ref<Page>()
let datetime: ComputedRef
let content = $ref('')

onMounted(async () => {
  page = await getPage(props.id)
  datetime = useDateFormat(page.created_ts, 'YYYY-MM-DD HH:mm:ss')
  content = marked.parse(page.content)
})
</script>

<template>
  <div v-if="!page">
    加载中...
  </div>
  <div v-else>
    <article text-base prose prose-coolgray dark:prose-white>
      <h1 text-3xl font-bold mb-0>
        {{ page.title }}
      </h1>
      <p text-md mt-1 text-coolgray>
        {{ page.author }} {{ datetime }}
      </p>
      <div v-html="content" />
    </article>
  </div>
</template>
