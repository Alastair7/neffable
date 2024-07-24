package com.example.neffable.home

import com.example.neffable.R
import androidx.compose.animation.core.animateFloatAsState
import androidx.compose.animation.core.tween
import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.material3.SnackbarHost
import androidx.compose.material3.SnackbarHostState
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.scale
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import kotlinx.coroutines.delay

@Composable
fun HomeScreen(viewModel: HomeViewModel, navController: NavController) {
    val snackbarMessage by remember { viewModel.snackbarMessage }
    val snackbarHostState = remember { SnackbarHostState() }

    var mainUserPulse by remember { mutableStateOf(false) }
    val mainUserScale by animateFloatAsState(
        targetValue = if (mainUserPulse) 1.2f else 1f,
        animationSpec = tween(durationMillis = 300)
    )

    var userConnectedPulse by remember { mutableStateOf(false) }
    val userConnectedScale by animateFloatAsState(
        targetValue = if (userConnectedPulse) 1.2f else 1f,
        animationSpec = tween(durationMillis = 300)
    )

    // Observa los cambios en userConnectedResponseReceived
    LaunchedEffect(viewModel.userConnectedResponseReceived.value) {
        if (viewModel.userConnectedResponseReceived.value) {
            userConnectedPulse = true
            delay(300)
            userConnectedPulse = false
            viewModel.userConnectedResponseReceived.value = false // Resetear el estado después de la animación
        }
    }

    Box(modifier = Modifier.fillMaxSize()) {
        Column(
            modifier = Modifier.fillMaxSize(),
            verticalArrangement = Arrangement.SpaceBetween,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            // UserConnectedTo
            Image(
                painter = painterResource(id = R.drawable.he),
                contentDescription = "Top Icon",
                modifier = Modifier
                    .size(120.dp)
                    .padding(top = 24.dp)
                    .scale(userConnectedScale)
            )
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.Center
            ) {
                Image(
                    painter = painterResource(id = R.drawable.love),
                    contentDescription = "Center Left Icon",
                    modifier = Modifier
                        .size(90.dp)
                        .padding(top = 24.dp)
                        .clickable {
                            viewModel.sendLoveEmotion("love")
                            mainUserPulse = true
                        }
                )
                Image(
                    painter = painterResource(id = R.drawable.happy),
                    contentDescription = "Center Middle Icon",
                    modifier = Modifier
                        .size(120.dp)
                        .padding(start = 36.dp)
                        .clickable {
                            viewModel.sendHappyEmotion("happy")
                            mainUserPulse = true
                        }
                )
                Image(
                    painter = painterResource(id = R.drawable.sad),
                    contentDescription = "Center Right Icon",
                    modifier = Modifier
                        .size(110.dp)
                        .padding(start = 48.dp)
                        .clickable {
                            viewModel.sendSadEmotion("sad")
                            mainUserPulse = true
                        }
                )
            }

            // MainUser
            LaunchedEffect(mainUserPulse) {
                if (mainUserPulse) {
                    delay(300)
                    mainUserPulse = false
                }
            }

            Image(
                painter = painterResource(id = R.drawable.she),
                contentDescription = "Bottom Icon",
                modifier = Modifier
                    .padding(bottom = 24.dp)
                    .size(90.dp)
                    .scale(mainUserScale)
            )
        }

        // Botón de regreso en la parte inferior izquierda
        Image(
            painter = painterResource(id = R.drawable.back_icon),
            contentDescription = "Back Icon",
            modifier = Modifier
                .size(70.dp)
                .padding(16.dp)
                .align(Alignment.BottomStart)
                .clickable { navController.navigate("mainPage") }
        )

        // SnackbarHost
        SnackbarHost(
            hostState = snackbarHostState,
            modifier = Modifier.align(Alignment.BottomCenter)
        )
    }
}
