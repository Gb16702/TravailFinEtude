export interface RequestInterface {
  Body?: BodyInit;
}

export const MethodEnum = {
  POST: "POST",
  PUT: "PUT",
  DELETE: "DELETE",
} as const;

export type Method = typeof MethodEnum[keyof typeof MethodEnum]

export interface RequestWithMethod<T extends Method> extends RequestInterface {
  Method?: T;
}

/**
 * This function is used to execute a request
 *
 * @param request
 * @param method
 * @param body
 * @returns {Promise<Response>}
 */
const Exec: <T extends Method>(request: string, method: RequestWithMethod<T>, body: RequestInterface
) => Promise<Response> = async <T extends Method>( request: string, method: RequestWithMethod<T>, body: RequestInterface) => {
  let response;

  if(!method.Method) {
      response = await fetch(request)
  } else {
      response = await fetch(request, {
        method: method.Method,
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
    })
  }
  return response;
};

export default Exec;
