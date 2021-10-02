<template>
  <div>
    <div v-on:click="reset" class="cursor-pointer">
      <h1 class="text-xl">
        <span class="font-semibold text-blue-700">FS</span><span class="font-medium">explorer</span>
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