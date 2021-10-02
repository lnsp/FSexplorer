<template>
  <div
    class="fixed z-20 inset-0 overflow-y-auto"
    aria-labelledby="modal-title"
    role="dialog"
    aria-modal="true"
    tabindex="0"
    ref="dialog"
    @keydown.stop="handleKeyboard"
  >
    <div
      class="
        flex
        items-end
        justify-center
        min-h-screen
        pt-4
        px-4
        pb-20
        text-center
        sm:block
        sm:p-0
      "
    >
      <div
        class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"
        aria-hidden="true"
      ></div>
      <span
        class="hidden sm:inline-block sm:align-middle sm:h-screen"
        aria-hidden="true"
        >&#8203;</span
      >
      <div
        class="
          inline-block
          align-bottom
          bg-white
          rounded-lg
          text-left
          overflow-hidden
          shadow-xl
          transform
          transition-all
          sm:my-8
          sm:align-middle
          sm:max-w-7xl
          sm:w-full
        "
      >
        <div class="bg-white p-4 font-mono text-sm overflow-x-scroll"><pre>{{ content }}</pre></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      content: '',
    }
  },
  props: ['path', 'visible'],
  watch: {
    visible: {
      immediate: true,
      handler(newValue, oldValue) {
        if (newValue && !oldValue) {
          console.log('visible value changed')
          this.fetch(this.path)
        }
      },
    },
  },
  mounted() {
    this.focus()
  },
  methods: {
    focus() {
      this.$refs.dialog.focus()
    },
    close() {
      this.$emit('onclose')
    },
    handleKeyboard(event) {
      switch (event.keyCode) {
        case 27:
          this.close()
          break
        case 72:
          this.close()
          break
      }
    },
    async fetch(path) {
      console.log('fetching content')
      this.content = await this.$axios.$get('/view', { params: { q: path } })
    },
  },
}
</script>