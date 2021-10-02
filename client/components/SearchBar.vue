<template>
  <div class="w-full max-w-xl relative shadow-sm">
    <div
      class="
        absolute
        insert-y-0
        left-0
        pl-3
        flex
        items-center
        pointer-events-none
        h-full
      "
    >
      <outline-search-icon class="w-4 h-4 text-gray-400" />
    </div>
    <input
      type="text"
      class="
        w-full
        border
        text-sm text-gray-700
        focus:text-black
        border-gray-200
        focus:border-blue-700
        rounded
        pl-9
      "
      v-on:keyup.enter="search(path)"
      v-model="path"
      tabindex="1"
    />
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  data () {
    return {
      path: this.viewedPath,
    }
  },
  computed: {
    viewedPath() {
      return this.$store.state.files.path;
    }
  },
  watch: {
    viewedPath (newValue, oldValue) {
      if (!newValue.startsWith('/')) {
        this.path = '/' + newValue;
      } else {
        this.path = newValue;
      }
    }
  },
  methods: {
    ...mapActions({
      search: 'files/search'
    })
  }
}
</script>