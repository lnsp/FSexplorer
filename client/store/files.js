export const state = () => ({
  list: [],
  path: '',
  history: [],
})

export const mutations = {
  search(state, { files, path }) {
    if (path != '') {
      state.history.push({
        path: JSON.parse(JSON.stringify(state.path)),
        list: JSON.parse(JSON.stringify(state.list)),
      })
    }
    state.list = files
    state.path = path
  },
  history(state, { steps }) {
    if (state.history.length == 1) {
      state.list = state.history.list[0]
      state.path = state.history.path[0]
      return
    }
    let item = state.history[state.history.length - steps]
    state.history = state.history.slice(0, -1)
    state.list = item.list
    state.path = item.path
  },
}

export const actions = {
  async search({ commit }, path) {
    try {
      let files = await this.$axios.$get('/search', { params: { q: path } })
      commit('search', { files, path })
    } catch {}
  },
  async history({ commit }, steps) {
    commit('history', { steps })
  },
}
