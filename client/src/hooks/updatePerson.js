import { notification } from 'antd'
import PERSON_API from "../api/personsAPI.JS"

const updatedPerson = async (request) => {
  const result = await PERSON_API()
    .put(`person`, request)
    .catch((error) => {
      return error.response ? error.response : error
    })
  if (result.status !== 200) {
    notification.error({
      message: 'Error',
      description: result.data.message,
      placement: 'bottomRight',
    })
    return {
      success: false,
    }
  }

  return result.data
}

export default updatedPerson