<template>
  <div
    class="p-4 focus:outline-none"
    @keydown="handleKeyboard"
    tabindex="0"
    ref="fileList"
  >
    <ul class="grid grid-cols-1 grid-rows-1 gap-3">
      <li
        v-for="(file, index) in files"
        v-bind:key="file.name"
        @click="selectAndNavigate(index)"
        class="
          file-list-item
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
          hover:bg-blue-100
          h-12
        "
        :class="{
          'border-blue-700': chosen(index),
          'bg-blue-100': chosen(index),
          'shadow-lg': chosen(index),
        }"
      >
        <div
          class="
            flex
            items-center
            justify-center
            text-gray-400
            w-10
            group-hover:text-blue-700
          "
          :class="{ 'text-blue-700': chosen(index) }"
        >
          <template v-if="file.isDirectory">
            <outline-folder-icon class="w-5 h-5" />
          </template>
          <template v-else>
            <outline-document-icon class="w-5 h-5" />
          </template>
        </div>
        <div class="w-full flex flex-row justify-between items-center">
          <div class="text-sm">{{ file.name }}</div>
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
    <FileModal
      v-if="isFileModalVisible"
      :visible="isFileModalVisible"
      :path="fileModalPath"
      @onclose="closeFileModal"
    />
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
      isFileModalVisible: false,
      fileModalPath: '',
    }
  },
  watch: {
    files(newValue, oldValue) {
      this.selected = 0
    },
  },
  mounted() {
    this.focus()
  },
  methods: {
    ...mapActions({
      select: 'files/search',
    }),
    focus() {
      this.$refs.fileList.focus()
    },
    chosen(index) {
      return this.selected == index
    },
    navigateBack() {
      this.$store.dispatch('files/history', 1)
    },
    navigateForward() {
      let file = this.files[this.selected]
      if (file.isDirectory) {
        this.select(file.path)
      } else {
        this.fileModalPath = file.path
        this.isFileModalVisible = true
      }
    },
    closeFileModal() {
      this.isFileModalVisible = false
      this.focus()
    },
    selectAndNavigate(index) {
      this.selected = index
      this.navigateForward()
    },
    handleKeyboard(event) {
      switch (event.keyCode) {
        case 72: // back
          this.navigateBack()
          break
        case 74: // down
          this.selected = Math.min(this.selected + 1, this.files.length - 1)
          this.$el.querySelector(`.file-list-item:nth-child(${this.selected+1})`).scrollIntoView(false)
          break
        case 75: // up
          this.selected = Math.max(this.selected - 1, 0)
          this.$el.querySelector(`.file-list-item:nth-child(${this.selected+1})`).scrollIntoView(false)
          break
        case 76: // forward
          this.navigateForward()
          break
      }
    },
  },
}
</script>