<script setup lang="ts">
import { defineProps } from 'vue';

defineProps({
  id: {
    type: String,
    required: true
  },
  label: {
    type: String,
    required: true
  },
  value: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: ''
  },
  type: {
    type: String,
    default: 'text'
  },
  error: {
    type: String,
    default: ''
  }
});
</script>
<template>
  <div class="mb-4">
    <label v-if="label" class="block text-sm font-medium text-gray-700">{{ label }}</label>
    <input
      :id="id"
      :value="value"
      :type="type"
      class="mt-1 px-4 block w-full border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500 outline-red-300 min-h-8"
    >
    <p v-if="error" class="text-red-500 text-sm mt-1">{{ error }}</p>
  </div>
</template>

