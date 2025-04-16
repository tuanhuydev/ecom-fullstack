<script setup lang="ts">

  // Declare the emit function and the event
  const emit = defineEmits<{
    (event: 'on-change-search', value: string): void
  }>()
  // hooks
  const store = useCartStore()

  // Ref
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
      <a href="/">
        <img src="/logo.svg" alt="logo" class="w-auto h-[18px]" >
      </a>

        <form class="flex-grow flex items-center gap-[14px] bg-white text-black  px-4 rounded-[8px] " @keyup="enterToSearchProduct" @submit="clickToSearchProduct">
          <input v-model="search" type="text" placeholder="Search" class="w-full h-[40px] focus:outline-none">
          <Icon name="uil:search" class="text-primary text-3xl" />
        </form>
        <div class="flex gap-[14px] items-center">
          <a class="relative block" href="/cart">
            <Icon name="uil:shopping-cart" class="text-red-50 text-3xl" />
            <span v-if="store.cartCount > 0" class="absolute bg-red-700 text-white rounded-full w-[18px] h-[18px] flex items-center justify-center text-sm font-bold -top-2 -right-2">
              {{ store.cartCount }}
            </span>
          </a>
          <a class="capitalize" href="/auth/sign-up">sign up</a>
        </div>
    </div>
  </header>
</template>

