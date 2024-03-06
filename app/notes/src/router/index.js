import { createRouter, createWebHistory } from 'vue-router'
import NoteEditorView from '../views/NoteEditorView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'editor',
      component: NoteEditorView
    }
  ]
})

export default router
