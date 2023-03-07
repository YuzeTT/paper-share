<script setup lang="ts" generic="T extends any, O extends any">
import getPages from '~/api/page_list'
import post from '~/api/post'
import type { Page } from '~/types/page.interface'

defineOptions({
  name: 'IndexPage',
})
let pages = $ref<Page[] | null>()
let msg = $ref<Page | null>()
const from = $ref({
  title: '',
  author: '',
  content: '',
})

onMounted(async () => {
  pages = await getPages()
})

const submit = async () => {
  msg = await post(from.title, from.author, from.content)
  pages?.unshift(msg)
  from.title = ''
  from.author = ''
  from.content = ''
}
</script>

<template>
  <div mb-2 border-2 border-coolgray-200 dark:border-zinc-700 rounded-lg bg-white dark:bg-zinc-800>
    <div p-2>
      <input v-model="from.title" type="text" placeholder="标题" bg-transparent text-2xl mb-1 outline-none w-full font-bold>
      <input v-model="from.author" type="text" placeholder="作者" bg-transparent text-md mb-1 outline-none w-full>
      <textarea id="" v-model="from.content" name="content" placeholder="记录点什么..." rows="3" w-full outline-none bg-transparent resize-none overf overflow-y-auto overflow-x-none text-lg />
    </div>
    <div flex justify-between p-2>
      <div flex items-center my-auto px-1>
        <button icon-btn text-lg>
          <div i-ri-file-add-line />
        </button>
      </div>
      <button flex items-center bg-emerald-500 text-white px-3 py-1 rounded-md @click="submit()">
        发布 <div i-ri-send-plane-2-line ml-1 />
      </button>
    </div>
  </div>
  <Transition mode="out-in">
    <div v-if="!pages" flex my-5>
      <div m-auto text-center>
        <div text-xl i-ri-loader-line animate-spin mx-auto mb-2 />
        <div>数据装载中</div>
      </div>
    </div>
    <div v-else space-y-2>
      <div v-for="(item, i) in pages" :key="i" bg-white dark:bg-zinc-800 rounded-lg p-4 cursor-pointer hover:shadow-md transition @click="$router.push(`/${item.id}`)">
        <h1 text-xl font-bold mb-0>
          {{ item.title || '无标题' }}
        </h1>
        <p text-md mt-1 text-coolgray>
          {{ useDateFormat(item.created_ts, 'YYYY-MM-DD HH:mm:ss').value }} | {{ item.author || '匿名' }}
        </p>
      </div>
    </div>
  </Transition>
</template>
