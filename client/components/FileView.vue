<template>
  <div class="p-4">
    <ul class="grid grid-cols-1 grid-rows-1 auto-rows-fr gap-4">
      <li
        v-for="file in files"
        v-bind:key="file.name"
        v-on:click="select(file.name)"
        class="
          flex flex-row
          items-center
          bg-white
          px-2
          py-2
          rounded
          border-2
          border-white
          shadow-sm
          group
          hover:shadow-lg
          hover:border-blue-700
        "
      >
        <div class="flex items-center justify-center text-gray-400 w-10">
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
import { mapActions } from 'vuex';

export default {
  computed: {
    files() {
      return this.$store.state.files.list
    },
  },
  methods: {
    ...mapActions({
      select: 'query/join',
    })
  }
}
</script>