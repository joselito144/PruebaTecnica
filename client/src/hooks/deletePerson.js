import PERSON_API from "../api/personsAPI.JS"


const deletePersons = async () => {
  const result = PERSON_API()
    .get(`person`)
    .catch((error) => {
      return error.response ? error.response : error
    })
  if (result.status !== 200) {
    return {}
  }
  return result.data.data
}

export default deletePersons