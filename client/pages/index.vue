<template>
<div>
  <FileView />
</div>
</template>

<script>
export default {
  mounted () {
    this.$store.dispatch('files/search', '/');
  },
}
</script>