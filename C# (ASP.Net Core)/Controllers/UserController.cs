using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Http;
using WebCRUD.Repositories;
using WebCRUD.DTOs;
using WebCRUD.Models;
using Dumpify;
using System;

namespace WebCRUD.Controllers;

public class UserController : Controller
{
	private readonly UserRepository _userRepository;

	// Inject the UserRepository through constructor injection
	public UserController(UserRepository userRepository )
	{
		_userRepository = userRepository;
	}

	[HttpGet]
	public IActionResult Index()
	{
		return LocalRedirect( "/users" );
	}

	[HttpGet]
	public async Task<IActionResult> List()
	{
		var users = await _userRepository.GetAll();
		string flashMessage = HttpContext.Session.GetString( "flash_message" );
		HttpContext.Session.Remove( "flash_message" );

		ViewBag.Flash = flashMessage;
		return View( users );
	}

	[HttpGet]
	public IActionResult New()
	{
		return View();
	}

	[HttpGet]
	public async Task<IActionResult> Edit( int Id )
	{
		var user = await _userRepository.GetByID( Id );
		return View( user );
	}

	[HttpGet]
	public async Task<IActionResult> Delete( int Id )
	{
		var user = await _userRepository.GetByID( Id );
		return View( user );
	}

	[HttpPost]
	public async Task<IActionResult> New(NewUserRequest request)
	{
		if ( !ModelState.IsValid ) return BadRequest( ModelState );

		// Insert form data into 'users' table
			var user = new UserModel
			{
				Name = request.Name,
				Email = request.Email,
				Website = request.Website
			};
			var insertID = await _userRepository.Insert( user );

		// Set flash message then redirect to 'users' table
			if( insertID <= 0 ){
				HttpContext.Session.SetString( "flash_message", "<div class='alert alert-danger'>Failed to add new user!</div>" );
				return LocalRedirect( "/users" );
			}
			HttpContext.Session.SetString( "flash_message", "<div class='alert alert-success'>The new user has been added successfully!</div>" );
			return LocalRedirect( "/users" );
	}

	[HttpPost]
	public async Task<IActionResult> Edit(EditUserRequest request)
	{
		if ( !ModelState.IsValid ) return BadRequest( ModelState );

		// Insert form data into 'users' table
			var user = new UserModel
			{
				Id = request.Id,
				Name = request.Name,
				Email = request.Email,
				Website = request.Website
			};
			var affectedRows = await _userRepository.Update( user );

		// Set flash message then redirect to 'users' table
			if( affectedRows <=0 ){
				HttpContext.Session.SetString( "flash_message", "<div class='alert alert-danger'>Failed to update user record!</div>" );
				return LocalRedirect( "/users" );
			}
			HttpContext.Session.SetString( "flash_message", "<div class='alert alert-success'>The user record has been updated successfully!</div>" );
			return LocalRedirect( "/users" );
	}
}