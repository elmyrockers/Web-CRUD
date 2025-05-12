using Microsoft.AspNetCore.Mvc;

namespace C___ASP.Net_Core_.Controllers;

public class UserController : Controller
{
	[HttpGet]
	public IActionResult Index()
	{
		return Redirect( "/users" );
	}

	[HttpGet]
	public IActionResult List()
	{
		return View();
	}
}