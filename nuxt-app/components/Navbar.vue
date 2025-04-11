<script setup lang="ts">
  // Declare the emit function and the event
  const emit = defineEmits<{
    (event: 'on-change-search', value: string): void
  }>()

  const search = ref('')

  const handleSearch = async (e: Event) => {
    e.preventDefault()
    emit('on-change-search', search.value)
  }

  const clickToSearchProduct = (e: Event) => {
    e.preventDefault()
    return handleSearch(e)
  }
  
  const enterToSearchProduct = (e: KeyboardEvent) => {
    if (e.key === 'Enter') {
      e.preventDefault()
      return handleSearch(e)
    }
  }

const emitDebouncedSearch = debounce((value: string) => {
    emit('on-change-search', value);
}, 300);

watch(search, (newValue) => {
    emitDebouncedSearch(newValue);
});
</script>
<template>
  <header class=" items-center p-4 bg-red-500 text-white">
    <div class="container mx-auto flex justify-between items-center gap-[14px]">
        <img src="/logo.svg" alt="logo" class="w-auto h-[18px]" >
        <form class="flex-grow flex items-center gap-[14px]" @keyup="enterToSearchProduct" @submit="clickToSearchProduct">
          <input v-model="search" type="text" placeholder="Search" class="w-full h-[40px] px-4 rounded-[8px] bg-white text-black">
          <button>hello</button>
        </form>
        <div class="flex gap-[14px]">
          <div class="cart">cart</div>
          <div class="signup">sign up</div>
        </div>
    </div>
  </header>
</template>

