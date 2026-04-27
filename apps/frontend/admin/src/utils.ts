export function getFieldError(errors: Array<unknown>) {
  const error = errors[0];

  if (typeof error === "string") {
    return error;
  }

  if (error instanceof Error) {
    return error.message;
  }

  return undefined;
}
