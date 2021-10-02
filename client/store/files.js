export const state = () => ({
  list: [],
})

export const mutations = {
  search(state, files) {
    state.list = files
    console.log('update files')
  },
}

export const actions = {
  async search({ commit, state }, query) {
    try {
      let files = await this.$axios.$get('/search', { params: { q: query } })
      commit('search', files)
    } catch {}
  },
}
