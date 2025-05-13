using Microsoft.AspNetCore.Mvc;
using WebCRUD.Repositories;
using Dumpify;

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
		return Redirect( "/users" );
	}

	[HttpGet]
	public IActionResult List()
	{
		var users = _userRepository.GetAll();
		users.Dump();
		return View();
	}

	[HttpGet]
	public IActionResult New()
	{
		return View();
	}

	[HttpGet]
	public IActionResult Edit()
	{
		return View();
	}

	[HttpGet]
	public IActionResult Delete()
	{
		return View();
	}
}