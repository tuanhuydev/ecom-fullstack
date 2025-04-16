<script setup lang="ts">
import type { Product } from '~/types/product'

// Apply layout
  definePageMeta({
  layout: 'standard'
})


const products = ref<Product[]>([])
const params = reactive({
  search: ''
})

const changeSearch = (newSearch: string) => {
  params.search = newSearch
  fetchProducts()
}

const fetchProducts = async () => {
    try {
      const { data, error }: { data: Product[] | null; error: unknown } = await $fetch('http://localhost:8080/products', { params })
      if (error || !data) {
        console.error('Error fetching products:', (error as {value: unknown }).value)
        return
      }
      products.value = data
    } catch (err) {
      console.error('Unexpected error:', err)
    }
}

// Lifecycle
onMounted(() => {
  fetchProducts()
})

</script>
<template>
  <div>
    <Navbar @on-change-search="changeSearch" />
    <div class="container mx-auto flex-grow py-16">
      <h1 class="text-3xl font-bold mb-3">Fresh Products</h1>
      <div>
        <div v-if="products.length === 0">
          <p>No products found.</p>
        </div>
        <div v-else class="flex flex-wrap gap-6">
          <ProductItem
            v-for="product in products"
            :key="product.id"
            :product="product"

          />
        </div>
      </div>
    </div>
  </div>
</template>