<template>
  <div>
    <div v-on:click="reset" class="cursor-pointer">
      <h1 class="text-xl">
        <span class="font-medium text-blue-600">fs</span><span class="text-gray-600">explorer</span>
      </h1>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    reset() {
      this.$store.dispatch('files/search', '/');
      this.$router.push('/');
    }
  }
}
</script>