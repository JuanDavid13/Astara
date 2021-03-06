//group them
/**
 * Función auxiliar que controlar los mensajes de error.
 *
 * @function
 * @param { string } errorName - Nombre del error
 * @returns { string } Devuelve el mensaje de error
*/
export function GetErrMsg(errorName) {
  switch(errorName) {
    case 'lackInput':{ return "Debe de completar todos los campos obligatorios"; }
    case 'nameRequired': { return "Debe de introducir un nombre"; }

    case 'noUser':{ return "El usuario introducido no existe"; }
    case 'alreadyCreated':{ return "Ya existe un usuario con ese nombre o email"; }
    case 'shortUserName':{ return "el nombre de usuario es demasiado corto"; }
    case 'emptyUser':{ return "debe de introducir un nombre de usuario"; }

    case 'missPass':{ return "La contraseña introducida es incorrecta"; }
    case 'diffPass': { return "Las contraseñas no coinciden"; }
    case 'emptyPass':{ return "Debes de introducir una contraseña"; }
    case 'wrongPass':{ return "Ha ocurrido un error al intentar cambiar la contraseña"; }
      
    case 'wrongEmail':{ return "el correo electrónico introducido no es valido. Ej: ejemplo@gmail.com"; }

    case 'wrongDate':{ return "La fecha introducida no es valida"; }

    case 'taskNoDeleted': { return "No se ha podido eliminar la tarea"; }

    case 'updateErr':{ return "Ha ocurrido un error al intentar actualizar la información"; }
    default:{ return "Ha ocurrido un error"; }
  }
}
