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
  <div>
    <Transition>
      <div v-if="!page" absolute top-0 left-0 h-100vh w-100vw bg-coolgray-100 dark:bg-zinc-900 flex z-100>
        <div m-auto text-center>
          <div text-xl i-ri-loader-line animate-spin mx-auto mb-2 />
          <div>数据装载中</div>
        </div>
      </div>
      <div v-else pt-6>
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
    </Transition>
  </div>
</template>
