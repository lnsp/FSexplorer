<template>
  <div class="p-4" @keydown="handleKeyboard" tabindex="0" ref="fileList">
    <ul class="grid grid-cols-1 grid-rows-1 auto-rows-fr gap-4">
      <li
        v-for="(file, index) in files"
        v-bind:key="file.name"
        @click="select(file.path)"
        class="
          flex flex-row
          items-center
          bg-white
          px-2
          py-2
          rounded
          border-2 border-white
          shadow-sm
          group
          hover:shadow-lg
          hover:border-blue-700
        "
        :class="{'border-blue-700': chosen(index)}"
        ref="files"
      >
        <div class="flex items-center justify-center text-gray-400 w-10 group-hover:text-blue-700" :class="{'border-blue-700': chosen(index) }">
          <template v-if="file.isDirectory">
            <outline-folder-icon class="w-6 h-6" />
          </template>
          <template v-else>
            <outline-document-icon class="w-6 h-6" />
          </template>
        </div>
        <div class="w-full flex flex-row justify-between items-center">
          <div>{{ file.name }}</div>
          <div class="flex flex-col items-end">
            <div v-if="!file.isDirectory" class="text-xs text-gray-500">
              {{ file.size }} bytes
            </div>
            <div v-if="file.mime" class="text-xs text-gray-500">
              {{ file.mime }}
            </div>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  computed: {
    files() {
      return this.$store.state.files.list
    },
  },
  data() {
    return {
      selected: 0,
    }
  },
  watch: {
    files(newValue, oldValue) {
      this.selected = 0
    },
  },
  mounted () {
    focus()
  },
  methods: {
    ...mapActions({
      select: 'files/search',
    }),
    focus () {
      this.$refs.fileList.$el.focus()
    },
    chosen (index) {
      return this.selected == index
    },
    navigateBack() {
      this.$store.dispatch('files/history', 1)
    },
    navigateForward() {
      this.select(this.files[this.selected].path)
    },
    handleKeyboard(event) {
      switch (event.keyCode) {
        case 72: // back
          this.navigateBack()
          break
        case 74: // down
          this.selected = (this.selected + 1) % this.files.length
          break
        case 75: // up
          this.selected =
            (this.selected + this.files.length - 1) % this.files.length
          break
        case 76: // forward
          this.navigateForward()
          break
      }
    },
  },
}
</script>