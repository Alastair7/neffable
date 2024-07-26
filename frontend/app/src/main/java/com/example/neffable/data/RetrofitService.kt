package com.example.neffable.data

import com.example.neffable.data.neffableModel.ApiResponse
import retrofit2.Response
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.http.GET

interface RetrofitService {
    @GET("api/test/ping")
    suspend fun ping(): Response<ApiResponse>
}

object RetrofitServiceFactory {
    fun makeRetrofitService(): RetrofitService {
        return Retrofit.Builder()
            .baseUrl("http://localhost:5001/")
            .addConverterFactory(GsonConverterFactory.create())
            .build()
            .create(RetrofitService::class.java)
    }
}

suspend fun fetchPingPong(): ApiResponse? {
    val service = RetrofitServiceFactory.makeRetrofitService()
    return try {
        println(service.ping().toString())
        val response = service.ping()
        if (response.isSuccessful) {
            println(response.body())
            response.body()
        } else {
            println("Error: ${response.code()} - ${response.message()}")
            null
        }
    } catch (e: Exception) {
        println("Exception: ${e.message}")
        null
    }
}

//interface RetrofitService {
//    @GET("v1/public/characters/{id}")
//    suspend fun getHero(
//        @Path("id") id: String,
//        @Query("apikey") apiKey: String,
//        @Query("ts") timestamp: String,
//        @Query("hash") hash: String
//    ): Response<RemoteResponse>
//}
//
//// RetrofitServiceFactory.kt
//object RetrofitServiceFactory {
//    fun makeRetrofitService(): RetrofitService {
//        return Retrofit.Builder()
//            .baseUrl("https://gateway.marvel.com:443/")
//            .addConverterFactory(GsonConverterFactory.create())
//            .build()
//            .create(RetrofitService::class.java)
//    }
//}
//
//fun getTimeStamp(): String {
//    return (System.currentTimeMillis() / 1000).toString()
//}
//
//fun md5(input: String): String {
//    val md = MessageDigest.getInstance("MD5")
//    val digest = md.digest(input.toByteArray())
//    val bigInt = BigInteger(1, digest)
//    return bigInt.toString(16).padStart(32, '0')
//}
//
//// Llamada a la API
//suspend fun fetchHero(id: String, apiKey: String, privateKey: String) {
//    val service = RetrofitServiceFactory.makeRetrofitService()
//    val timestamp = getTimeStamp()
//    val hash = md5(timestamp + privateKey + apiKey)
//
//    try {
//        val response = service.getHero(id, apiKey, timestamp, hash)
//        if (response.isSuccessful) {
//
//            val heroResponse = response.body()
//            println("SUCCESS -- LOADING DATA")
//
//            if (heroResponse != null) {
//                val characterName = heroResponse.data.results[0].name
//                println("Character Name: $characterName")
//            }
//
//            println(heroResponse)
//
//        } else {
//            println("Error: ${response.code()} - ${response.message()}")
//
//            when (response.code()) {
//                401 -> println("Unauthorized: Check your API key.")
//                403 -> println("Forbidden: You don't have access to this resource.")
//                404 -> println("Not Found: The resource does not exist.")
//                409 -> println("Conflict: There is a conflict with the current state of the resource.")
//                else -> println("Unknown error occurred.")
//            }
//        }
//    } catch (e: Exception) {
//        println("Exception: ${e.message}")
//    }
//}