// 状態管理の設定 zastan

import { create } from 'zustand'

type EditedTask = {
    id: number
    title: string
}

// 返り値がない場合はvoid型にする
type State = {
    editedTask: EditedTask
    updateEditedTask: (payload: EditedTask) => void
    resetEditedTask: () => void
}

const useStore = create<State>((set) => ({
    editedTask: { id: 0, title: ''},
    updateEditedTask: (payload) =>
        set({
            editedTask: payload,
        }),
    resetEditedTask: () => set({ editedTask: { id: 0, title: ''}}),
}))

export default useStore