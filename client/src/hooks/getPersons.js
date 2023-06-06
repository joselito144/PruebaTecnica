import PERSONAPI from "../api/personsAPI.js"


const getPersons = async () => {
  const result = await PERSONAPI()
    .get(`person`)
    .catch((error) => {
      return error.response ? error.response : error
    })

  console.log(result)
  if (result.status !== 200) {
    return {}
  }
  return result.data.data
}

export default getPersons