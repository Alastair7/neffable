package com.example.neffable

import MainPage
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.example.neffable.home.HomeScreen
import com.example.neffable.home.HomeViewModel

@Composable
fun NavGraph(modifier: Modifier = Modifier, navController: NavHostController =
    rememberNavController()) {
    val homeViewModel: HomeViewModel = viewModel()
    NavHost(navController = navController, startDestination = "mainPage") {
        composable("mainPage") { MainPage(navController) }
        composable("homeScreen") { HomeScreen(viewModel = homeViewModel,
            navController = navController) }
    }
}
