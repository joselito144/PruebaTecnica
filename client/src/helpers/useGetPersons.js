import getPersons from "../hooks/getPersons";
import { useEffect, useState } from "react";

const useGetPersons = () => {
  const [data, setData] = useState();
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    getPersons().then(
      response => {
        setData(response);
        setIsLoading(false);
    })
  }, [])

  return {
    data, isLoading
  }
}

export default useGetPersons;