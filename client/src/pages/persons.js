import { Table, Spin, Tooltip, Button } from 'antd'
import { EditOutlined, DeleteOutlined } from '@ant-design/icons'
import useGetPersons from '../helpers/useGetPersons'


function Persons() {

  const { data: persons, isLoading } = useGetPersons()

  const updatedPerson = (record) => {
    console.log(record)
  }

  const deletePerson = (record) => {
    console.log(record)
  }

  if (!isLoading) {
    console.log(persons)
  }

  const columns = [
    {
      title: 'Cédula',
      dataIndex: 'id',
      key: 'id',
      align: 'center',
    },
    {
      title: 'Nombre',
      dataIndex: 'name',
      key: 'name',
      align: 'center',
    },
    {
      title: 'Apellido',
      dataIndex: 'lastName',
      key: 'lastName',
      align: 'center',
    },

    {
      title: 'Opciones',
      render: (_, record) => {
        return (
          <>
            <Tooltip title='Actualizar'>
              <Button
                onClick={() => updatedPerson(record)}
                icon={<EditOutlined />}
                shape='round'
                size='large'
              />
            </Tooltip>
            <Tooltip title='Eliminar'>
              <Button
                onClick={() => deletePerson(record)}
                icon={<DeleteOutlined />}
                shape='round'
                size='large'
              />
            </Tooltip>
          </>



        )


      }
    }
  ]
  return (
    <>
      {
        isLoading ?
          <Spin />
          :
          <>
            <p>Personas</p>
            <Table
              dataSource={persons}
              columns={columns}
            />
          </>

      }
    </>
  )
}

export default Persons