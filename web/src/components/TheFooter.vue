<template>
  <nav text-xl mt-6 inline-flex gap-2>
    <button icon-btn @click="toggleDark()">
      <div dark:i-ri-moon-line i-ri-sun-line />
    </button>

    <a
      icon-btn i-ri-github-line
      rel="noreferrer"
      href="https://github.com/YuzeTT"
      target="_blank"
      title="GitHub"
    />
  </nav>
</template>
