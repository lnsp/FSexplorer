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
    <iframe class="hidden" ref="download" />
    <div
      class="
        flex
        justify-center
        min-h-screen
        p-4
        text-center
      "
    >
      <div
        class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"
        aria-hidden="true"
      ></div>
      <div
        class="
          flex
          align-bottom
          bg-white
          rounded-lg
          text-left
          shadow-xl
          overflow-hidden
          transform
          transition-all
          sm:w-full
          sm:my-8
          sm:align-middle
        "
      >
        <div class="bg-white p-4 text-md flex flex-col justify-center items-center w-full" v-if="error">
          <div>
            <div class="text-center uppercase text-xs font-medium text-gray-500">status</div>
            <div class="text-3xl font-bold text-red-600">{{ error.status }}</div>
          </div>
          <div class="mt-4 text-red-600">{{ error.message }}</div>
        </div>
        <div class="bg-white p-4 font-mono text-sm overflow-x-scroll" v-else>
          <pre>{{ content }}</pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      error: null,
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
          this.view(this.path)
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
        case 68:
          this.download(this.path)
          break
      }
    },
    async download(path) {
      console.log('attempting download')
      console.log(this.$refs.download)
      this.$refs.download.src =
        this.$axios.defaults.baseURL +
        '/view?download=true&q=' +
        encodeURIComponent(path)
    },
    async view(path) {
      try {
        this.content = await this.$axios.$get('/view', { params: { q: path } })
      } catch (err) {
        this.error = {
          status: err.response.status,
          message: err.response.data,
        }
      }
    },
  },
}
</script>