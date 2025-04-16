<script setup lang="ts">
import { ref } from 'vue';
import InputField from '~/components/InputField.vue';

const form = ref({
  name: '',
  email: '',
  password: '',
  confirmPassword: ''
});

const errors = ref({
  name: '',
  email: '',
  password: '',
  confirmPassword: ''
});

const handleSubmit = async () => {
  console.log('Form submitted:', form.value);
};
</script>

<template>
  <div class="flex justify-center items-center h-screen bg-gray-100">
    <div class="w-full max-w-md bg-white p-8 rounded shadow">
      <h1 class="text-2xl font-bold mb-6">Sign Up</h1>

      <form @submit.prevent="handleSubmit">
        <InputField
          id="name"
          v-model="form.name"
          label="Name"
          :error="errors.name"
          type="text"
        />

        <InputField
          id="email"
          v-model="form.email"
          label="Email"
          :error="errors.email"
          type="email"
        />

        <InputField
          id="password"
          v-model="form.password"
          label="Password"
          :error="errors.password"
          type="password"
        />

        <InputField
          id="confirmPassword"
          v-model="form.confirmPassword"
          label="Confirm Password"
          :error="errors.confirmPassword"
          type="password"
        />

        <button
          type="submit"
          class="w-full bg-red-500 text-white py-2 px-4 rounded hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
        >
          Sign Up
        </button>
      </form>

      <p class="mt-4 text-center text-sm text-gray-600">
        Already have an account? <a href="/sign-in" class="text-red-500 hover:underline">Sign In</a>
      </p>
      <p class="mt-2 text-center text-sm text-gray-600">
        <a href="/" class="text-blue-500 hover:underline">Back to Home</a>
      </p>
    </div>
  </div>
</template>