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
      v-on:keyup.enter="search(query)"
      v-model="query"
    />
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  data () {
    return {
      query: this.$store.state.query.path,
    }
  },
  computed: {
    viewedQuery() {
      return this.$store.state.query.path
    }
  },
  watch: {
    viewedQuery (newValue, oldValue) {
      this.query = newValue;
    }
  },
  methods: {
    ...mapActions({
      search: 'query/update'
    })
  }
}
</script>