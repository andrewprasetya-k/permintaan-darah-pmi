import { inject, provide, ref, type Component, type Ref } from 'vue'

export interface PageHeaderAction {
  label: string
  to?: string
  icon?: Component
  onClick?: () => void
  variant?: 'primary' | 'secondary' | 'danger'
}

export function usePageHeader() {
  const actions = ref<PageHeaderAction[]>([])

  const setActions = (newActions: PageHeaderAction[]) => {
    actions.value = newActions
  }

  const clearActions = () => {
    actions.value = []
  }

  provide('pageHeaderActions', {
    actions: actions as Ref<PageHeaderAction[]>,
    setActions,
    clearActions,
  })

  return {
    actions,
    setActions,
    clearActions,
  }
}

export function useSetPageHeaderActions() {
  const pageHeaderContext = inject<{
    actions: Ref<PageHeaderAction[]>
    setActions: (actions: PageHeaderAction[]) => void
    clearActions: () => void
  }>('pageHeaderActions', {
    actions: ref<PageHeaderAction[]>([]),
    setActions: () => {},
    clearActions: () => {},
  })

  return {
    setActions: pageHeaderContext.setActions,
    clearActions: pageHeaderContext.clearActions,
  }
}
