<script setup lang="ts">
import { defineProps } from 'vue';
import { useField } from 'vee-validate';
const props = defineProps({
  name: {
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
const { value, errorMessage } = useField(props.name, {
  validateOnMount: true,
  validateOnBlur: true,
  validateOnChange: true
});

</script>
<template>
  <div class="mb-4">
    <label v-if="label" class="block text-sm font-medium text-gray-700">{{ label }}</label>
    <input
      v-model="value"
      :placeholder="placeholder"
      :type="type"
      class="mt-1 px-4 block w-full border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500 outline-red-300 min-h-8"
    >
    <p v-if="errorMessage" class="text-red-500 text-sm mt-1">{{errorMessage}}</p>
  </div>
</template>

