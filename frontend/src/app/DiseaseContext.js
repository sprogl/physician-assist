import { createContext, useReducer } from "react"

export const DiseaseContext = createContext()

export const sympReducer = (state, action) => {
  switch (action.type) {
    case 'SET_SYMPTOMS':
      return {
        symptoms: action.payload
      }
    case 'CREATE_SYMPTOM':
      return {
        symptoms: [action.payload, ...state.symptoms]
      }
    default:
      return state
  }
}

export const DiseaseContextProvider = ({ children }) => {
  const [state, dispatch] = useReducer(sympReducer, {
    symptoms: null
  })

  dispatch({ type: ''})
  return (
    <DiseaseContext.Provider value={{ ...state, dispatch }}>
      { children }
    </DiseaseContext.Provider>
  )
}