package proxy_address_forwarding_demo

import javax.ws.rs.GET
import javax.ws.rs.Path
import javax.ws.rs.Produces
import javax.ws.rs.core.Response
import javax.servlet.http.HttpServletRequest
import javax.ws.rs.core.Context
import javax.ws.rs.core.UriInfo
import javax.ws.rs.core.MultivaluedMap


data class HttpServletRequestData(
    val requestURI: String?,
    val remoteAddr: String?,
    val remoteHost: String?,
    val remotePort: Int?,
    val localAddr: String?,
    val localName: String?,
    val localPort: Int?,
    val contextPath: String?,
    val pathInfo: String?,
    val requestURL: String?,
    val parameterMap: MutableMap<String, Array<String>>?,
    val serverName: String?
)

data class UriInfoData(
    val requestUri: String?,
    val absolutePath: String?,
    val path: String?,
    val baseUri: String?,
    val queryParameters: MultivaluedMap<String, String>?
)

data class ResponseData(val servletRequest: HttpServletRequestData, val uriInfo: UriInfoData?) {

}

@Path("/request-info")
public open class Resource {

    @Context
    private var servletRequest: HttpServletRequest?  = null
    @Context
    private var uriInfo: UriInfo? = null

    @Path("/")
    @GET
    @Produces("application/json")
    open fun printRequestInfo():  Response {
        val response = ResponseData(
            HttpServletRequestData(
                servletRequest?.requestURI,
                servletRequest?.remoteAddr,
                servletRequest?.remoteHost,
                servletRequest?.remotePort,
                servletRequest?.localAddr,
                servletRequest?.localName,
                servletRequest?.localPort,
                servletRequest?.contextPath,
                servletRequest?.pathInfo,
                servletRequest?.requestURL.toString(),
                servletRequest?.parameterMap,
                servletRequest?.serverName
            ),
            UriInfoData(
                uriInfo?.requestUri?.toString(),
                uriInfo?.absolutePath.toString(),
                uriInfo?.path,
                uriInfo?.baseUri.toString(),
                uriInfo?.queryParameters))

        return Response.ok(response).build()
    }
}

