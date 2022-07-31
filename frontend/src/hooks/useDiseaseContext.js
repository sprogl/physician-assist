import { DiseaseContext } from "../app/DiseaseContext"
import { useContext } from "react"

export const useDiseaseContext = () => {
  const context = useContext(DiseaseContext)

  if (!context) {
    throw Error('useDiseaseContext must be used inside a DiseaseContextProvider')
  }

  return context
}