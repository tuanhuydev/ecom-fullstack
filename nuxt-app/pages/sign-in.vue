<script setup lang="ts">

const form = reactive({
  email: '',
  password: ''
});


const handleSubmit = async () => {
  console.log('Form submitted:', form);
};
</script>

<template>
  <div class="flex justify-center items-center h-screen bg-gray-100">
    <div class="w-full max-w-md bg-white p-8 rounded shadow">
      <h1 class="text-2xl font-bold mb-6">Sign In</h1>

      <form @submit.prevent="handleSubmit">
        <div class="mb-4">
          <InputField
            id="email"
            v-model="form.email"
            label="Email"
            type="email"
          />
        </div>

        <div class="mb-4">
          <InputField
            id="password"
            v-model="form.password"
            label="Password"
            type="Password"
          />
        </div>

        <button
          type="submit"
          class="w-full bg-red-500 text-white py-2 px-4 rounded hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
        >
          Sign In
        </button>
      </form>

      <p class="mt-4 text-center text-sm text-gray-600">
        Don't have an account? <a href="/sign-up" class="text-red-500 hover:underline">Sign Up</a>
      </p>
      <p class="mt-2 text-center text-sm text-gray-600">
        <a href="/" class="text-blue-500 hover:underline">Back to Home</a>
      </p>
    </div>
  </div>
</template>