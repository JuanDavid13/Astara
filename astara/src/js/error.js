//group them
export function GetErrMsg(errorName) {
  switch(errorName) {
    case 'lackInput':{ return "Debe de completar todos los campos obligatorios"; }

    case 'noUser':{ return "El usuario introducido no existe"; }
    case 'alreadyCreated':{ return "Ya existe un usuario con ese nombre o email"; }
    case 'shortUserName':{ return "el nombre de usuario es demasiado corto"; }
    case 'emptyUser':{ return "debe de introducir un nombre de usuario"; }

    case 'missPass':{ return "La contraseña introducida es incorrecta"; }
    case 'diffPass': { return "Las contraseñas no coinciden"; }
    case 'emptyPass':{ return "Debes de introducir una contraseña"; }
      
    case 'wrongEmail':{ return "el correo electrónico introducido no es valido. Ej: ejemplo@gmail.com"; }

    case 'updateErr':{ return "Ha ocurrido un error al intentar actualizar la información"; }
    default:{ return "Ha ocurrido un error"; }
  }
}
