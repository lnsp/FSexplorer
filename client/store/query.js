export const state = () => ({
  path: '/'
})

export const mutations = {
  update(state, path) {
    state.path = path;
  },

  join(state, path) {
    if (!state.path.endsWith('/')) {
      state.path += '/';
    }
    state.path += path;
  }
}

export const actions = {
  update({ commit, state, dispatch }, path) {
    commit('update', path);
    dispatch('files/search', state.path, { root: true });
  },

  join({ commit, state, dispatch }, path) {
    commit('join', path);
    dispatch('files/search', state.path, { root: true });
  }
}