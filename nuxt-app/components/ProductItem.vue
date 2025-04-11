<script setup lang="ts">
import { defineProps } from 'vue';
import type { Product } from '~/types/product';
  defineProps({
    product: {
      type: Object as () => Product,
      required: true,
    }
  })

</script>
<template>
  <div class="p-3 bg-white rounded-lg shadow-md w-fit flex flex-col gap-3">
    <div class="w-[184px] h-[184px] bg-gray-200 rounded-lg mb-2">
      <img
      v-if="product.imageUrl"
        :src="product.imageUrl"
        alt="Product Image"
        class="w-full h-full object-cover rounded-lg"
      >

    </div>
    <div class="flex-col gap-2">
      <h3 class="text-xl font-bold">{{ product.name }}</h3>
      <h4>${{ product.price }}</h4>
    </div>
    <button class="bg-transparent text-red-500 border-red-500 border rounded-sm p-2">
      <span class="text-sm">Add to Cart</span>
    </button>
  </div>
</template>