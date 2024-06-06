package com.example.neffable.home
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.launch

class HomeViewModel : ViewModel() {
    // TODO: Añadir la lógica aquí
    val snackbarMessage = mutableStateOf("")

    fun sendMessage(message: String) {
        viewModelScope.launch {
            snackbarMessage.value = message
        }
    }
}
